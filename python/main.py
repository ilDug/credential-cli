#! /usr/bin/env python3

from typing import Annotated
import typer
from core.config import ROOT, SECRET
from rich import print
from rich.panel import Panel
from rich.prompt import Prompt
from pathlib import Path
import re
import secrets
import string
import yaml


app = typer.Typer()


@app.command("mongo")
def mongo_utils(
    root: Annotated[bool, typer.Option("--root", "-r", help="Root user")] = False,
    output: Annotated[
        str,
        typer.Option(
            "--out", "-o", help="location of secrets folder", rich_help_panel="Commons"
        ),
    ] = None,
):
    # set the destination folder for the secrets
    dest = SECRET if output is None else Path(output)
    if not dest.exists():
        print(
            f"[yellow]Warning:[/yellow] {dest} [white]does not exist. Creating secrets folder..."
        )
        dest.mkdir(parents=True)

    print(Panel("[white][bold]MongoDB Utilities[/bold]\ncreating user credentials"))
    username = typer.prompt("\nEnter username") if not root else "root"
    print(f"Username: {username}")
    host = Prompt.ask(
        "\nEnter the [blue]hostname:port[/blue] of database (e.g. [purple]mongo1.dag.lan:27017,mongo2.dag.lan:27017,mongo3.dag.lan:27017[/purple]) "
    )

    # check whether the host follow the correct regular expression
    pattern = re.compile(r"([a-z0-9\.\-]+:\d+)+")
    if not (match := pattern.findall(host)):
        print("[red]Error: Invalid hostname:port format")
        raise typer.Exit(code=1)

    is_replica = len(match) > 1
    if is_replica:
        print(
            "[blue]info:[white] Multiple hosts detected. Using replica set configuration"
        )
        rs = typer.prompt("\nEnter the replica set name", default="rs0")

    database = typer.prompt("\nEnter the database name")

    overwrite = True
    pw_file = dest / f"MONGO_{username.upper()}_PW"

    if pw_file.exists():
        print(f"[yellow]Warning:[/yellow] {pw_file} [white]already exists.")
        overwrite = typer.confirm("Do you want to overwrite the file?")

    if overwrite:
        length = 64
        characters = string.ascii_letters + string.digits + "-_"
        password = "".join(secrets.choice(characters) for _ in range(length))
        pw_file.write_text(password)
        print(f"[green]Password file created at {pw_file}")

    # retrive the content of the password file
    pw = pw_file.read_text()
    if is_replica:
        mongo_cs = f"mongodb://{username}:{pw}@{host}/{database}?authSource=admin&replicaSet={rs}"
    else:
        mongo_cs = f"mongodb://{username}:{pw}@{host}/{database}?authSource=admin"

    l = f"""
    [b]Username:[/b] [green]{username}[/green]
    [b]Password:[/b] [white]{pw}[/white]
    [b]Database:[/b] [yellow]{database}[/yellow]
    [b]Host:[/b] [blue]{host}[/blue]
    [b]Connection String:[/b] [red]{mongo_cs}[/red]
    """
    print(Panel(l, title="MongoDB Credentials", style="white"))

    # Create YAML file with the credentials data
    credentials_data = {
        "username": username,
        "password": pw,
        "database": database,
        "host": host,
        "connection_string": mongo_cs,
    }
    yaml_file = dest / f"mongo_{username.lower()}_credentials.yaml"
    with yaml_file.open("w") as file:
        yaml.dump(credentials_data, file)
    print(f"[green]YAML file created at {yaml_file}")

    # name = typer.prompt("[grey]Enter yours name")
    # print(Panel(f"[red]Hello, {name}!", title="Welcome"), style="bold green")


@app.callback()
def callback():
    pass

    # typer.secho(f"Welcome here", fg=typer.colors.WHITE)
    # print(f"[white]Output Secret folder:[/white] { output }")


if __name__ == "__main__":
    app()

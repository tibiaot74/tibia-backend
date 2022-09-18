import os

import mysql.connector

from behave.runner import Context
from behave.model import Scenario
from typing import List

def before_all(context: Context):
    context.url = "http://api:7474/api"

    def query_db(query: str) -> List[tuple]:
        connection = mysql.connector.connect(
            host=os.getenv("DB_HOST"),
            user=os.getenv("DB_USER"),
            password=os.getenv("DB_PASSWORD"),
            database=os.getenv("DB_NAME"),
            port=os.getenv("DB_PORT")
        )
        cursor = connection.cursor(dictionary=True)
        cursor.execute(query)
        result = cursor.fetchall()
        connection.commit()
        connection.close()
        return result

    context.query_db = query_db


def after_scenario(context: Context, scenario: Scenario):
    result = context.query_db("SHOW TABLES")
    for row in result:
        for _, table in row.items():
            context.query_db(f"DELETE FROM {table}")

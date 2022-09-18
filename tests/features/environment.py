import os

import mysql.connector


def before_all(context):
    context.url = "http://api:7474/api"

    def query_db(query):
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


def after_scenario(context, scenario):
    tables = [table for tables in context.query_db(
        "SHOW TABLES") for table in tables.values()]
    for table in tables:
        context.query_db(f"DELETE FROM {table}")

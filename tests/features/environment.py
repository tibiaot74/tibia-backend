import mysql.connector


def before_all(context):
    context.url = "http://localhost:7474/api"

    def query_db(query):
        connection = mysql.connector.connect(host="localhost", user="root", password="YES", database="tibia")
        cursor = connection.cursor(dictionary=True)
        cursor.execute(query)
        result = cursor.fetchall()
        connection.commit()
        connection.close()
        return result

    context.query_db = query_db


def after_scenario(context, scenario):
    tables = [table for tables in context.query_db("SHOW TABLES") for table in tables.values()]
    for table in tables:
        context.query_db(f"DELETE FROM {table}")
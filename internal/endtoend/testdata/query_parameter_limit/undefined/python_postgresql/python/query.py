# Code generated by sqlc. DO NOT EDIT.
# versions:
#   sqlc v1.17.0
# source: query.sql
import sqlalchemy
import sqlalchemy.ext.asyncio

from querytest import models


DELETE_BAR_BY_ID = """-- name: delete_bar_by_id \\:execrows
DELETE FROM bar WHERE id = :p1
"""


DELETE_BAR_BY_ID_AND_NAME = """-- name: delete_bar_by_id_and_name \\:execrows
DELETE FROM bar
WHERE id = :p1
AND name1 = :p2
AND name2 = :p3
AND name3 = :p4
"""


class Querier:
    def __init__(self, conn: sqlalchemy.engine.Connection):
        self._conn = conn

    def delete_bar_by_id(self, *, id: int) -> int:
        result = self._conn.execute(sqlalchemy.text(DELETE_BAR_BY_ID), {"p1": id})
        return result.rowcount

    def delete_bar_by_id_and_name(self, *, id: int, name1: str, name2: str, name3: str) -> int:
        result = self._conn.execute(sqlalchemy.text(DELETE_BAR_BY_ID_AND_NAME), {
            "p1": id,
            "p2": name1,
            "p3": name2,
            "p4": name3,
        })
        return result.rowcount


class AsyncQuerier:
    def __init__(self, conn: sqlalchemy.ext.asyncio.AsyncConnection):
        self._conn = conn

    async def delete_bar_by_id(self, *, id: int) -> int:
        result = await self._conn.execute(sqlalchemy.text(DELETE_BAR_BY_ID), {"p1": id})
        return result.rowcount

    async def delete_bar_by_id_and_name(self, *, id: int, name1: str, name2: str, name3: str) -> int:
        result = await self._conn.execute(sqlalchemy.text(DELETE_BAR_BY_ID_AND_NAME), {
            "p1": id,
            "p2": name1,
            "p3": name2,
            "p4": name3,
        })
        return result.rowcount

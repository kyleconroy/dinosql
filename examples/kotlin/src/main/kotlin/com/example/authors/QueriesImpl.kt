// Code generated by sqlc. DO NOT EDIT.

package com.example.authors

import java.sql.Connection
import java.sql.SQLException

import sqlc.runtime.ExecuteQuery
import sqlc.runtime.ListQuery
import sqlc.runtime.RowQuery

const val createAuthor = """-- name: createAuthor :one
INSERT INTO authors (
          name, bio
) VALUES (
  ?, ?
)
RETURNING id, name, bio
"""

const val deleteAuthor = """-- name: deleteAuthor :exec
DELETE FROM authors
WHERE id = ?
"""

const val getAuthor = """-- name: getAuthor :one
SELECT id, name, bio FROM authors
WHERE id = ? LIMIT 1
"""

const val listAuthors = """-- name: listAuthors :many
SELECT id, name, bio FROM authors
ORDER BY name
"""

class QueriesImpl(private val conn: Connection) : Queries {

  @Throws(SQLException::class)
  override fun createAuthor(name: String, bio: String?): RowQuery<Author> {
    return object : RowQuery<Author>() {
      override fun execute(): Author {
        return conn.prepareStatement(createAuthor).use { stmt ->
          this.statement = stmt
          stmt.setString(1, name)
          stmt.setString(2, bio)

          val results = stmt.executeQuery()
          if (!results.next()) {
            throw SQLException("no rows in result set")
          }
          val ret = Author(
                results.getLong(1),
                results.getString(2),
                results.getString(3)
            )
          if (results.next()) {
              throw SQLException("expected one row in result set, but got many")
          }
          ret
        }
      }
    }
  }

  @Throws(SQLException::class)
  override fun deleteAuthor(id: Long): ExecuteQuery {
    return object : ExecuteQuery() {
      override fun execute() {
        conn.prepareStatement(deleteAuthor).use { stmt ->
          this.statement = stmt
          stmt.setLong(1, id)

          stmt.execute()
        }
      }
    }
  }

  @Throws(SQLException::class)
  override fun getAuthor(id: Long): RowQuery<Author> {
    return object : RowQuery<Author>() {
      override fun execute(): Author {
        return conn.prepareStatement(getAuthor).use { stmt ->
          this.statement = stmt
          stmt.setLong(1, id)

          val results = stmt.executeQuery()
          if (!results.next()) {
            throw SQLException("no rows in result set")
          }
          val ret = Author(
                results.getLong(1),
                results.getString(2),
                results.getString(3)
            )
          if (results.next()) {
              throw SQLException("expected one row in result set, but got many")
          }
          ret
        }
      }
    }
  }

  @Throws(SQLException::class)
  override fun listAuthors(): ListQuery<Author> {
    return object : ListQuery<Author>() {
      override fun execute(): List<Author> {
        return conn.prepareStatement(listAuthors).use { stmt ->
          this.statement = stmt
          
          val results = stmt.executeQuery()
          val ret = mutableListOf<Author>()
          while (results.next()) {
              ret.add(Author(
                results.getLong(1),
                results.getString(2),
                results.getString(3)
            ))
          }
          ret
        }
      }
    }
  }

}


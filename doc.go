// Package sqlarfs provides an implementation of [io/fs.FS]
// for [SQLite Archive Files] via [database/sql].
//
// [SQLite Archive Files]: https://sqlite.org/sqlar.html
//
// Recall that the schema of an SQLite archive file is:
//   CREATE TABLE sqlar(
//   name TEXT PRIMARY KEY,  -- name of the file (incl. dirs; dirs w trlg "/" ?)
//   mode INT,  -- access permissions
//   mtime INT, -- last modif'cn time
//   sz INT,    -- original file size
//   data BLOB  -- compressed content
//  );
package sqlarfs


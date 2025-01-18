# FILE I/O CODE BANK
# with open("<file_name>", "<mode>") as file: # w for write (replaces all content), r for read, a for append (adds to end of file rather than replacing)
#   file.write("<stuff>") # for writing
#   content = file.read() # for reading

# ***********************************************************************
# PYTHON sqlite3 CODE BANK

# CREATING A DB
# conn = sqlite3.connect('<db_name>')

# ***********************************************************************
# GENERALLY EXECUTING A COMMAND
# conn.execute(<sql>, (argument1, argument2)) # note: use ? for placeholders in your SQL. If you only have one argument, keep a trailing comma! (I don't know why!)
# Example:
# conn.execute("SELECT * FROM students WHERE id = ?", (id,))
# conn.commit()
# or
# cursor = conn.cursor()
# cursor.execute("SELECT * FROM students WHERE id = ?", (id,)")
# conn.commit()

# ***********************************************************************
# FOR QUERIES
# cursor = conn.cursor() # you need to use cursor for these
# cursor.execute(<sql>, (argument1, argument2))
# <thing> = cursor.fetchone() # when only one row is expected e.g. your query included WHERE id = ____
# <thing_id> = thing[0]
# <thing_column2_val> = thing[1]

# or if you're expecting more than one row (e.g. "SELECT * FROM students" with no `WHERE`)
# <things> = cursor.fetchall()
# <thing1> = <things>[0]
# <thing1_id> = <thing1>[0]

# ***********************************************************************
# SQL CODE BANK
# <placeholder> means you should replace this with something!

# CREATING A TABLE
# CREATE TABLE IF NOT EXISTS <table_name>
# (id INTEGER PRIMARY KEY,
# <column_name2> <data_type2>,
# <column_name3> <data_type3>)

# QUERYING THE DB
# SELECT * from <table_name> WHERE <column_name> = <value>
# Example:
# SELECT * from students WHERE id = ?

# DELETING A ROW
# DELETE from <table_name> WHERE <column_name> = <value>

# ***********************************************************************
# EXAMPLE FUNCTIONS
# def create_connection(db_name: str)->sqlite3.Connection:
#     conn = sqlite3.connect(db_name)
#     return conn

# def create_table(conn: sqlite3.Connection):
#     sql = """
#     CREATE TABLE BLAH BLAH BLAH

# """
#     conn.execute(sql)
#     conn.commit()

# def query_student(conn: sqlite3.Connection, param1, param2):
#     sql = """
#     SELECT * from students BLAH BLAH BLAH

# """
#     cursor = conn.cursor()
#     cursor.execute(sql, (param1, param2))
#     conn.commit()
#     student = cursor.fetchone()
#     return student

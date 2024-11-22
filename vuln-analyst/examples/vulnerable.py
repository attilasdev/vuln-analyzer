# Test file with various vulnerabilities

import sqlite3
import hashlib

# SQL Injection vulnerability (HIGH)
def get_user(user_id):
    conn = sqlite3.connect('database.db')
    cursor = conn.cursor()
    query = "SELECT * FROM users WHERE id = " + user_id  # Unsafe
    cursor.execute(query)
    return cursor.fetchone()

# Hardcoded credentials (HIGH)
DATABASE_PASSWORD = "db_password_123"
API_SECRET = "my_secret_key_123"

# Weak cryptography (HIGH)
def hash_password(password):
    return hashlib.md5(password.encode()).hexdigest()  # MD5 is weak

# Command injection (HIGH)
def execute_command(cmd):
    import os
    os.system("echo " + cmd)  # Unsafe
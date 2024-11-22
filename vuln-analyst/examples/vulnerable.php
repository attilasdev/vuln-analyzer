<?php
// Test file with various vulnerabilities

// SQL Injection (HIGH)
$user_id = $_GET['id'];
$query = "SELECT * FROM users WHERE id = " . $user_id;  // Unsafe
$result = mysql_query($query);

// Path traversal (HIGH)
$file = $_GET['file'];
include($file);  // Unsafe

// Hardcoded credentials (HIGH)
define('DB_PASSWORD', 'super_secret_123');
$apiKey = "1234567890abcdef";

// Command injection (HIGH)
$command = $_POST['command'];
system("ping " . $command);  // Unsafe
?>
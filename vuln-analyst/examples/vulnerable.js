// Test file with various vulnerabilities

// Hardcoded credentials (HIGH)
const API_KEY = "1234567890abcdef";
const SECRET_TOKEN = "super_secret_token_123";

// XSS Vulnerability (HIGH)
function displayUser(userData) {
    document.getElementById('user').innerHTML = userData.name;
    document.write(userData.description);
}

// Weak password validation (HIGH)
function validatePassword(password) {
    return password.length >= 6;  // Too short
}

// Unsafe eval usage (HIGH)
function processData(data) {
    eval('console.log(' + data + ')');
}
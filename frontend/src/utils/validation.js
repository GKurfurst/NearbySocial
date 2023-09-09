
// 验证用户名是否至少为5个字符
export function validateUsername(username) {
    return username.length >= 3;
}
export function validateTelephone(telephone) {
    return telephone.length >= 11;
}

// 验证密码是否至少为6个字符
export function validatePassword(password) {
    return password.length >= 6;
}

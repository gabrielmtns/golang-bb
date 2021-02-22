
const LOCAL_STORAGE_NAME = 'bb-teste'

export class Token {
  static setToken(user) {
    localStorage.setItem(LOCAL_STORAGE_NAME, btoa(JSON.stringify(user)))
  }

  static getToken() {
    const token = localStorage.getItem(LOCAL_STORAGE_NAME)
    if (!token) {
      return null
    }
    return JSON.parse(atob(token)).Token
  }
}
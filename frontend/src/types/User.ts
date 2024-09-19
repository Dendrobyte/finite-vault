export type User = {
    username: string
    email: string
    balance: number
    auth_token: string
  }

export type ValidatedUser = {
  valid: boolean
  email: string
}

export type UserTransaction = {
  creation_ts: number
  amount: number
  description: string
}
export interface ErrorResponse {
  message: string
}

export interface User {
  id: number
  uuid: string
  first_name: string
  last_name: string
  email: string
  created_at: string
  updated_at: string
}

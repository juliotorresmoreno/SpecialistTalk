export type IUser = {
  id: number
  email: string
  username: string
  name: string
  lastname: string
  date_birth?: string
  imgSrc?: string
  country?: string
  nationality?: string
  facebook?: string
  linkedin?: string
}
export class User {
  id: number
  email: string
  username: string
  name: string
  lastname: string
  date_birth?: string
  imgSrc?: string
  country?: string
  nationality?: string
  facebook?: string
  linkedin?: string

  constructor(data: IUser) {
    this.id = data.id
    this.email = data.email
    this.username = data.username
    this.name = data.name
    this.lastname = data.lastname
    this.date_birth = data.date_birth

    this.imgSrc = data.imgSrc
    this.country = data.country
    this.nationality = data.nationality
    this.facebook = data.facebook
    this.linkedin = data.linkedin
  }

  getFullName() {
    return [this.name, this.lastname].join(' ')
  }
}

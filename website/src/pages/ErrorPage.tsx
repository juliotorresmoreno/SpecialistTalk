import React from 'react'
import Header from '../components/Header'
import { HTTPError } from '../types/http'

type ErrorPageProps = {
  error?: Error | HTTPError
}

const ErrorPage: React.FC<ErrorPageProps> = ({ error }) => {
  const header = {
    title: 'Error ',
    description: 'programa de super poderes',
  }
  if (!error) return null

  return (
    <>
      <Header {...header} />
      <main>{error.message}</main>
    </>
  )
}

export default ErrorPage

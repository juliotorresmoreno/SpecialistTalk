import React from 'react'
import { HTTPError } from '../types/http'

type ErrorPageProps = {
  error?: Error | HTTPError
}

const ErrorPage: React.FC<ErrorPageProps> = ({ error }) => {
  if (!error) return null

  return (
    <>
      <header>
        <h1>Error</h1>  
      </header>
      <main>{error.message}</main>
    </>
  )
}

export default ErrorPage

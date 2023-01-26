import React from 'react'
import Header from '../components/Header'

const SignInPage: React.FC = () => {
  const header = {
    title: 'SignIn',
    description: 'programa de super poderes',
  }

  return (
    <>
      <Header {...header} />
      <main>SignIn</main>
    </>
  )
}

export default SignInPage

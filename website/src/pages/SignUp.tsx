import React from 'react'
import Header from '../components/Header'

const SignUpPage: React.FC = () => {
  const header = {
    title: 'SignUp',
    description: 'programa de super poderes',
  }

  return (
    <>
      <Header {...header} />
      <main>SignUp</main>
    </>
  )
}

export default SignUpPage

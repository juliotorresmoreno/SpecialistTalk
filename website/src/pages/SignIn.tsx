import React from 'react'
import { Col, Row } from 'react-bootstrap'
import Header from '../components/Header'
import SingInForm from '../components/SignInForm'

const SignInPage: React.FC = () => {
  const header = {
    title: 'SignIn',
    description: 'programa de super poderes',
  }

  return (
    <>
      <Header {...header} />
      <main>
        <Row>
          <Col md={{ span: 6 }}></Col>
        </Row>
      </main>
    </>
  )
}

export default SignInPage

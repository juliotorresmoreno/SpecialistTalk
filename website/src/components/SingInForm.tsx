import React, { useState } from 'react'
import { Button, Form } from 'react-bootstrap'

type formData = {
  email: string
  password: string
}

type RequiredFieldsOnly<T> = {
  [Property in keyof T]?: T[Property]
}

type SingInFormProps = {
  onSubmit: (payload: formData) => void
  errors: RequiredFieldsOnly<formData> & {
    message?: string
  }
}

const SingInForm: React.FC<SingInFormProps> = ({ onSubmit, errors }) => {
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const defaultSubmit: React.FormEventHandler<HTMLFormElement> = (e) => {
    e.preventDefault()

    onSubmit({ email, password })
  }
  return (
    <Form onSubmit={defaultSubmit}>
      <Form.Group className="mb-3" controlId="formBasicEmail">
        <Form.Label>Email address</Form.Label>
        <Form.Control
          type="email"
          placeholder="Enter email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />
        {errors.email ? (
          <Form.Text className="text-muted">{errors.email}</Form.Text>
        ) : null}
      </Form.Group>

      <Form.Group className="mb-3" controlId="formBasicPassword">
        <Form.Label>Password</Form.Label>
        <Form.Control
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
        {errors.password ? (
          <Form.Text className="text-muted">{errors.password}</Form.Text>
        ) : null}
      </Form.Group>
      {errors.message ? (
        <Form.Text className="text-muted">{errors.message}</Form.Text>
      ) : null}

      <Button variant="primary" type="submit">
        Submit
      </Button>
    </Form>
  )
}

export default SingInForm

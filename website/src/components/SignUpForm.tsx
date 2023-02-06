import React, { useState } from 'react'
import { Button, Col, Form, Row } from 'react-bootstrap'
import config from '../config'
import withFormHandler from '../hoc/withFormHandler'

type formData = {
  username: string
  name: string
  lastname: string
  email: string
  password: string
  password_repeat: string
}

type RequiredFieldsOnly<T> = {
  [Property in keyof T]?: T[Property]
}

type SignUpFormProps = {
  isLoading?: boolean
  onSubmit: (payload: formData) => void
  errors: RequiredFieldsOnly<formData> & {
    message?: string
  }
}

const SignUpForm: React.FC<SignUpFormProps> = ({
  onSubmit,
  errors = {},
  isLoading = false,
}) => {
  const [username, setUsername] = useState('')
  const [name, setName] = useState('')
  const [lastname, setLastname] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [password_repeat, setPasswordRepeat] = useState('')
  const [terms_agreed, setTermsAgreed] = useState(false)
  const defaultSubmit: React.FormEventHandler<HTMLFormElement> = (e) => {
    e.preventDefault()

    onSubmit({
      username,
      name,
      lastname,
      email,
      password,
      password_repeat,
    })
  }
  if (isLoading) return <>Loading</>
  return (
    <Form onSubmit={defaultSubmit}>
      <Row>
        <Col>
          <Form.Group className="mb-3" controlId="form-name">
            <Form.Label>Name</Form.Label>
            <Form.Control
              type="text"
              name="name"
              placeholder="Name"
              value={name}
              autoComplete="first-name"
              onChange={(e) => setName(e.target.value)}
            />
            {errors.name ? (
              <Form.Text className="text-muted">{errors.name}</Form.Text>
            ) : null}
          </Form.Group>
        </Col>
        <Col>
          <Form.Group className="mb-3" controlId="form-lastname">
            <Form.Label>Last Name</Form.Label>
            <Form.Control
              type="text"
              name="lastname"
              placeholder="Last Name"
              value={lastname}
              autoComplete="last-name"
              onChange={(e) => setLastname(e.target.value)}
            />
            {errors.lastname ? (
              <Form.Text className="text-muted">{errors.lastname}</Form.Text>
            ) : null}
          </Form.Group>
        </Col>
      </Row>

      <Form.Group className="mb-3" controlId="form-username">
        <Form.Label>Username</Form.Label>
        <Form.Control
          type="text"
          name="username"
          placeholder="Username"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
        />
        {errors.username ? (
          <Form.Text className="text-muted">{errors.username}</Form.Text>
        ) : null}
      </Form.Group>

      <Form.Group className="mb-3" controlId="form-email">
        <Form.Label>Email address</Form.Label>
        <Form.Control
          type="email"
          name="email"
          placeholder="Enter email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          autoComplete="email"
        />
        {errors.email ? (
          <Form.Text className="text-muted">{errors.email}</Form.Text>
        ) : null}
      </Form.Group>

      <Row>
        <Col>
          <Form.Group className="mb-3" controlId="form-password">
            <Form.Label>Password</Form.Label>
            <Form.Control
              name="password"
              type="password"
              placeholder="Password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
            {errors.password ? (
              <Form.Text className="text-muted">{errors.password}</Form.Text>
            ) : null}
          </Form.Group>
        </Col>
        <Col>
          <Form.Group className="mb-3" controlId="form-password-repeat">
            <Form.Label>Repeat password</Form.Label>
            <Form.Control
              type="password"
              placeholder="Repeat password"
              value={password_repeat}
              onChange={(e) => setPasswordRepeat(e.target.value)}
            />
            {errors.password_repeat ? (
              <Form.Text className="text-muted">
                {errors.password_repeat}
              </Form.Text>
            ) : null}
          </Form.Group>
        </Col>
      </Row>

      {errors.message ? (
        <div>
          <Form.Text className="text-danger">
            <strong>{errors.message}</strong>
          </Form.Text>
        </div>
      ) : null}

      <Form.Group className="mb-3" controlId="form-terms-agreed">
        <Form.Check
          type="checkbox"
          label="Agree terms and conditions"
          checked={terms_agreed}
          onChange={() => setTermsAgreed(!terms_agreed)}
        />
      </Form.Group>

      <Button disabled={!terms_agreed} variant="primary" type="submit">
        Submit
      </Button>
    </Form>
  )
}

const url = config.baseUrl + '/auth/sing-up'

export default withFormHandler<SignUpFormProps>(SignUpForm, 'POST', url)

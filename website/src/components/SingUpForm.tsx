import React, { useState } from 'react'
import { Button, Form } from 'react-bootstrap'

type formData = {
  username: string
  name: string
  lastname: string
  email: string
  password: string
}

type RequiredFieldsOnly<T> = {
  [Property in keyof T]?: T[Property]
}

type SingUpFormProps = {
  onSubmit: (payload: formData) => void
  errors: RequiredFieldsOnly<formData> & {
    message?: string
  }
}

const SingUpForm: React.FC<SingUpFormProps> = ({ onSubmit, errors }) => {
  const [username, setUsername] = useState('')
  const [name, setName] = useState('')
  const [lastname, setLastname] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const defaultSubmit: React.FormEventHandler<HTMLFormElement> = (e) => {
    e.preventDefault()

    onSubmit({
      username,
      name,
      lastname,
      email,
      password,
    })
  }
  return (
    <Form onSubmit={defaultSubmit}>
      <Form.Group className="mb-3">
        <Form.Label>Name</Form.Label>
        <Form.Control
          type="text"
          placeholder="Name"
          value={name}
          onChange={(e) => setName(e.target.value)}
        />
        {errors.name ? (
          <Form.Text className="text-muted">{errors.name}</Form.Text>
        ) : null}
      </Form.Group>
      <Form.Group className="mb-3">
        <Form.Label>Last Name</Form.Label>
        <Form.Control
          type="text"
          placeholder="Last Name"
          value={lastname}
          onChange={(e) => setLastname(e.target.value)}
        />
        {errors.lastname ? (
          <Form.Text className="text-muted">{errors.lastname}</Form.Text>
        ) : null}
      </Form.Group>
      <Form.Group className="mb-3">
        <Form.Label>Username</Form.Label>
        <Form.Control
          type="text"
          placeholder="Last Name"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
        />
        {errors.username ? (
          <Form.Text className="text-muted">{errors.username}</Form.Text>
        ) : null}
      </Form.Group>

      <Form.Group className="mb-3">
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

      <Form.Group className="mb-3">
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

      <Form.Group className="mb-3">
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
      <Form.Group className="mb-3">
        <Form.Check type="checkbox" label="Check me out" />
      </Form.Group>
      <Button variant="primary" type="submit">
        Submit
      </Button>
    </Form>
  )
}

export default SingUpForm

import React from 'react'
import { Button, Form } from 'react-bootstrap'
import { Link } from 'react-router-dom'
import config from '../config'
import withFormHandler from '../hoc/withFormHandler'
import useFormValue from '../hooks/useFormValue'

type formData = {
  email: string
  password: string
}

type RequiredFieldsOnly<T> = {
  [Property in keyof T]?: T[Property]
}

type SignInFormProps = {
  isLoading?: boolean
  onSubmit: (payload: formData) => void
  errors: RequiredFieldsOnly<formData> & {
    message?: string
  }
}

const SignInForm: React.FC<SignInFormProps> = ({
  onSubmit,
  errors = {},
  isLoading = false,
}) => {
  const [email, setEmail] = useFormValue()
  const [password, setPassword] = useFormValue()
  const defaultSubmit: React.FormEventHandler<HTMLFormElement> = (e) => {
    e.preventDefault()

    onSubmit({ email, password })
  }
  if (isLoading) return <>Loading</>
  return (
    <Form onSubmit={defaultSubmit}>
      <Form.Group className="mb-3" controlId="form-email">
        <Form.Label>Email address</Form.Label>
        <Form.Control
          name="email"
          type="email"
          placeholder="Enter email"
          value={email}
          onChange={setEmail}
          autoComplete="email"
        />
        {errors.email ? (
          <Form.Text className="text-muted">{errors.email}</Form.Text>
        ) : null}
      </Form.Group>
      <Form.Group className="mb-3" controlId="form-password">
        <Form.Label>Password</Form.Label>
        <Form.Control
          type="password"
          placeholder="Password"
          value={password}
          onChange={setPassword}
        />
        {errors.password ? (
          <Form.Text className="text-muted">{errors.password}</Form.Text>
        ) : null}
      </Form.Group>
      {errors.message ? (
        <div>
          <Form.Text className="text-danger">
            <strong>{errors.message}</strong>
          </Form.Text>
        </div>
      ) : null}
      <Button variant="primary" type="submit">
        <span className="material-symbols-outlined">login</span>
        Submit
      </Button>
      &nbsp;&nbsp;&nbsp;
      <Link to="/recovery-password">Recuperar cuenta</Link>
    </Form>
  )
}

const url = config.baseUrl + '/auth/sing-in'

export default withFormHandler<SignInFormProps>(SignInForm, 'POST', url)

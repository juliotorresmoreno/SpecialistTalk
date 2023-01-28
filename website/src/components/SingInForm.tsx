import React from 'react'
import { Button, Form } from 'react-bootstrap'
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

type SingInFormProps = {
  isLoading?: boolean
  onSubmit: (payload: formData) => void
  errors: RequiredFieldsOnly<formData> & {
    message?: string
  }
}

const SingInForm: React.FC<SingInFormProps> = ({
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
          type="email"
          placeholder="Enter email"
          value={email}
          onChange={setEmail}
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
        <Form.Text className="text-muted">{errors.message}</Form.Text>
      ) : null}

      <Button variant="primary" type="submit">
        Submit
      </Button>
    </Form>
  )
}

const url = config.baseUrl + '/auth/sing-in'

export default withFormHandler<SingInFormProps>(SingInForm, 'POST', url)

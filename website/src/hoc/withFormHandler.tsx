import React, { useState } from 'react'

type Errors = {
  [x: string]: string
} & { message?: string }

type WithFormHandlerProps = {
  onSuccess?: (payload: any) => void
  onFailure?: (errors: Errors) => void
}

export default function withFormHandler<T = any>(
  WrappedComponent: React.ComponentType<any>,
  method: string,
  url: string
) {
  return function WithFormHandler({ onSuccess }: WithFormHandlerProps) {
    const [errors, setErrors] = useState<Errors>({})
    const [isLoading, setIsLoading] = useState(false)
    const onSubmit = async (payload: any) => {
      try {
        setIsLoading(true)
        const response = await fetch(url, {
          method,
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(payload),
        })
        setIsLoading(false)
        const body = await response.json()
        if (!response.ok) {
          const [key, value] = body.message.split(':')
          if (value !== undefined) {
            setErrors({ [key]: value })
          } else {
            setErrors({ message: key })
          }
          return
        }
        onSuccess && onSuccess(body)
      } catch (error: any) {
        if ('message' in error) setErrors({ message: error.message })
      }
    }
    const config: any = {
      isLoading,
      onSubmit,
      errors,
    }
    return <WrappedComponent {...config} />
  }
}

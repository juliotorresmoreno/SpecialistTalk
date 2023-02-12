import React from 'react'
import { Session } from '../models/session'
import { useApi } from '../services/api'

type Errors = {
  [x: string]: string
} & { message?: string }

type WithFormHandlerProps = {
  onSuccess?: (payload: Session) => void
  onFailure?: (errors: Errors) => void
}

export default function withFormHandler<T = any>(
  WrappedComponent: React.ComponentType<any>,
  method: string,
  url: string
) {
  return function WithFormHandler({ onSuccess }: WithFormHandlerProps) {
    const { error, isLoading, apply } = useApi<T, Session>(method, url)

    const onSubmit = async (payload: T) => {
      try {
        const response = await apply(payload)

        onSuccess && onSuccess(response)
      } catch (error: any) {}
    }
    const config: any = {
      isLoading,
      onSubmit,
      errors: error ?? {},
    }
    return <WrappedComponent {...config} />
  }
}

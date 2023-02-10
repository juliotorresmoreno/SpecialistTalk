import React, { useState } from 'react'
import { useAppSelector } from '../store/hooks'
import { HTTPError } from '../types/http'

type ApiOpts = {}

export function useGetData<Response = any>(url: string, opts?: ApiOpts) {
  const { error, isLoading, apply } = useApi<null, Response>('GET', url, opts)

  return { isLoading, error, get: () => apply(null) }
}

export function useApi<Payload = any, Response = any>(
  method: string,
  url: string,
  opts: ApiOpts = {}
) {
  const session = useAppSelector((state) => state.auth.session)
  const [isLoading, setIsLoading] = useState<boolean>(false)
  const [error, setError] = useState<HTTPError | null>(null)

  const apply = async (payload: Payload) => {
    setIsLoading(true)
    setError(null)

    try {
      const args: RequestInit = {
        method: method,
        headers: {
          'X-API-Key': session?.token ?? '',
          'Content-Type': 'application/json',
        },
      }
      if (payload) args.body = JSON.stringify(payload)

      const response = await fetch(url, args)
      const content = await response.json()
      if (!response.ok) {
        throw new Error(content.message)
      }
      setIsLoading(false)
      return content as Response
    } catch (err) {
      setIsLoading(false)
      setError({
        message: (err as Error).message,
      })
      throw err
    }
  }

  return { isLoading, error, apply }
}

export function useAdd<Payload = any, Response = any>(
  url: string,
  opts?: ApiOpts
) {
  const { error, isLoading, apply } = useApi<Payload, Response>(
    'POST',
    url,
    opts
  )

  return { isLoading, error, add: apply }
}

export function useUpdate<Payload = any, Response = any>(
  url: string,
  id: string,
  opts?: ApiOpts
) {
  const _url = url + '/' + id
  const { error, isLoading, apply } = useApi<Payload, Response>(
    'PATCH',
    _url,
    opts
  )

  return { isLoading, error, update: apply }
}

export function useRemove(url: string, id: string, opts?: ApiOpts) {
  const _url = url + '/' + id
  const { error, isLoading, apply } = useApi<null, void>('DELETE', _url, opts)

  return {
    isLoading,
    error,
    remove: async () => {
      await apply(null)
    },
  }
}

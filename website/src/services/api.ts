import React, { useEffect, useState } from 'react'
import { useAppSelector } from '../store/hooks'
import { HTTPError } from '../types/http'

type GetDataOpts = {
  onlyLogged?: boolean
}

export function useGetData(url: string, opts: GetDataOpts = {}) {
  const session = useAppSelector((state) => state.auth.session)
  const [isLoading, setIsLoading] = useState<boolean>(false)
  const [data, setData] = useState<any>(null)
  const [error, setError] = useState<HTTPError | null>(null)
  const { onlyLogged = false } = opts

  useEffect(() => {
    if (isLoading) return
    if (onlyLogged && !session?.token) return
    setIsLoading(true)
    setError(null)

    fetch(url, {
      headers: {
        'X-API-Key': session?.token ?? '',
      },
    })
      .then(async (response) => {
        const content = await response.json()
        if (!response.ok) {
          setError(content)
          return
        }
        setData(content)
      })
      .catch((err: Error) => setError({ message: err.message }))
      .finally(() => setIsLoading(false))
  }, [url])

  return { isLoading, error, data }
}

export function useApi<Payload = any, Response = any>(
  method: string,
  url: string
) {
  const session = useAppSelector((state) => state.auth.session)
  const [isLoading, setIsLoading] = useState<boolean>(false)
  const [error, setError] = useState<HTTPError | null>(null)

  const apply = async (payload: Payload) => {
    setIsLoading(true)
    setError(null)

    try {
      const response = await fetch(url, {
        method: method,
        headers: {
          'X-API-Key': session?.token ?? '',
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(payload),
      })
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

export function useAdd<Payload = any, Response = any>(url: string) {
  const { error, isLoading, apply } = useApi<Payload, Response>('POST', url)

  return { isLoading, error, add: apply }
}

export function useUpdate<Payload = any, Response = any>(
  url: string,
  id: string
) {
  const _url = url + '/' + id
  const { error, isLoading, apply } = useApi<Payload, Response>('POST', _url)

  return { isLoading, error, update: apply }
}

export function useRemove<Response = any>(url: string, id: string) {
  const _url = url + '/' + id
  const { error, isLoading, apply } = useApi<void, Response>('POST', _url)

  return { isLoading, error, remove: apply }
}

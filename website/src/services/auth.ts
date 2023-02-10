import React from 'react'
import config from '../config'
import { useGetData } from './api'

export function useGetSession() {
  const url = config.baseUrl + '/auth/session'
  const { data, error, isLoading } = useGetData(url, { onlyLogged: true })

  return { isLoading, error, session: data }
}

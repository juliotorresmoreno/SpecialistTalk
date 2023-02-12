import React from 'react'
import config from '../config'
import { IChat } from '../models/chat'
import * as api from './api'

export function useGet(id: string) {
  const url = config.baseUrl + '/messages/'
  const { get, error, isLoading } = api.useGetData(url + id)

  return { isLoading, error, get }
}

type Add = {
  code: string
  message: string
}

export function useAdd() {
  const url = config.baseUrl + '/messages'
  const { error, isLoading, add } = api.useAdd<Add, IChat>(url)

  return { isLoading, error, add }
}

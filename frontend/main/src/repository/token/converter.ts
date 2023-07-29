import { SetTokenReqData } from '@/api/main/apiClient'
import { Token } from '@/components/model/auth/type'

export const convertTokenToReqData = (token: Token): SetTokenReqData => {
  return { token }
}
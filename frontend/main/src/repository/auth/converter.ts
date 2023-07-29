import { SignupReqData, SignupResData } from '@/api/auth/apiClient'
import { SignupSeed, Token } from '@/components/model/auth/type'

export const convertSignupSeedToReqData = ({ email, password }: SignupSeed): SignupReqData => {
  return { email, password }
}

export const convertSignupResDataToToken = (resData: SignupResData): Token => {
  return resData.token
}
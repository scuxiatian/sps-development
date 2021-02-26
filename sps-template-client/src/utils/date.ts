import moment, { MomentInput } from 'moment'

export const formatTime = (time: MomentInput, format: string) => {
  return moment(time).format(format)
}

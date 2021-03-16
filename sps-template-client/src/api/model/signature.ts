export interface SignatureParams {
  name: string;
  password: string;
  url: string;
  description: string;
  isPublic: boolean;
  ownerId: string;
}

export interface ChangeSignaturePasswordParams {
  id: string;
  password: string;
  newPassword: string;
}

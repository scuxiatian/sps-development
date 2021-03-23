export interface SignatureParams {
  name: string;
  password: string;
  url: string;
  description: string;
  isPublic: boolean;
  ownerId: string;
}

export interface ChangeSignaturePasswordParams extends ValidateSignaturePasswordParams {
  newPassword: string;
}

export interface ValidateSignaturePasswordParams {
  id: number;
  password: string;
}

export interface UseSignatureParams {
  recordId?: number;
  signatureId?: number;
  description?: string;
}

export interface SignatureRecordParams {
  id: string;
  signatures: Array<UseSignatureParams>;
}

export interface AuthorityParams {
  authorityId: string;
  authorityName: string;
  parentId: string;
}

export interface CopyAuthorityParams {
  authority: AuthorityParams;
  oldAuthorityId: string;
}

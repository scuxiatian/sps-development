interface CasbinInfo {
  path: string;
  method: string;
}

export interface CasbinInReceiveParams {
  authorityId: string;
  casbinInfos?: Array<CasbinInfo>
}

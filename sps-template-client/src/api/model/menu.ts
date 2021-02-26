interface Meta {
  title: string;
  keepAlive?: boolean;
  defaultMenu?: boolean;
  icon?: string;
  hidden?: boolean;
}

export interface MenuParams {
  parentId?: string;
  path?: string;
  name: string;
  hidden?: boolean;
  component?: any;
  sort?: number;
  meta: Meta;
  children?: Array<MenuParams>;
  menuId?: string;
}

export interface MenuAuthorityParams {
  authorityId: string;
  menus: Array<MenuParams>
}

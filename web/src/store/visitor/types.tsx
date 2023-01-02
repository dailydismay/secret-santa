export interface TokensPair {
  accessToken: string | null;
  refreshToken: string | null;
}

export interface Visitor {
  id: string;
  firstName: string;
  lastName: string;
  authProviderID: string;
  avatarURL: string;
  createdAt: string;
}

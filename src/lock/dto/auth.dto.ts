export class AuthDto {
  access_token: string;

  device_id: string;

  device_secret: string;

  constructor({
    accessToken,
    deviceId,
    deviceSecret
  }) {
    this.access_token = accessToken;
    this.device_id = deviceId;
    this.device_secret = deviceSecret;

    return this;
  }
}
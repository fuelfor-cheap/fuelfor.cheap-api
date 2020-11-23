export class LockDto {
  longitude: number;

  latitude: number;

  access_token: string;

  device_id: string;

  device_secret: string;

  fuel_type: string;

  account_id: string;

  constructor({
    lon,
    lat,
    accessToken,
    deviceId,
    deviceSecret,
    fuelType,
    accountId
  }) {
    this.longitude = lon;
    this.latitude = lat;
    this.access_token = accessToken;
    this.device_id = deviceId;
    this.device_secret = deviceSecret;
    this.fuel_type = fuelType;
    this.account_id = accountId;

    return this;
  }
}
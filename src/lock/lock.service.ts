import { Injectable } from '@nestjs/common';
import { invalidSchema } from 'src/utils/invalidschema';

import { LockDto } from './dto/lock.dto';

@Injectable()
export class LockService {
  async createOne(payload: LockDto): Promise<any> {
    const {
      longitude,
      latitude
    } = payload;
    if (longitude === 0 || latitude === 0) {
      // return error showing an invalid schema
      return invalidSchema(payload);
    }

    
  }
}

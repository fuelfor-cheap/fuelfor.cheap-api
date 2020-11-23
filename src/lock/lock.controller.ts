import {
  Controller,
  Post,
  Body
} from '@nestjs/common';

import { LockService } from './lock.service';
import { LockDto } from './dto/lock.dto';
import { AuthDto } from './dto/auth.dto';

@Controller('/lock')
export class LockController {
  constructor(private readonly lockService: LockService) {}

  @Post('/')
  createNewLock(@Body() lockDto: LockDto): Promise<any> {
    return this.lockService.createOne(lockDto);
  }

  @Post('/existing')
  getOne(@Body() authDto: AuthDto): Promise<any> {

  }
}

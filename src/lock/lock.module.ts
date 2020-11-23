import { Module } from '@nestjs/common';
import { LockController } from './lock.controller';
import { LockService } from './lock.service';

@Module({
  imports: [],
  controllers: [LockController],
  providers: [LockService],
})
export class LockModule {}

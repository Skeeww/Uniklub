import { Module } from '@nestjs/common';
import { TypeMaterielsService } from './type_materiels.service';
import { TypeMaterielsController } from './type_materiels.controller';
import { ClubsModule } from '../clubs.module';

@Module({
  controllers: [TypeMaterielsController],
  providers: [TypeMaterielsService],
  imports: [ClubsModule],
})
export class TypeMaterielsModule {}

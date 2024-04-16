import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { ClubsModule } from './clubs/clubs.module';
import { TypeMaterielsModule } from './clubs/type_materiels/type_materiels.module';

@Module({
  imports: [ClubsModule, TypeMaterielsModule],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}

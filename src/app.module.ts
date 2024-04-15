import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { ClubsModule } from './clubs/clubs.module';

@Module({
  imports: [ClubsModule],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}

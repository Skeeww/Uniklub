import {
  Body,
  Controller,
  Delete,
  Get,
  HttpException,
  HttpStatus,
  Param,
  Patch,
  Post,
} from '@nestjs/common';
import { TypeMaterielsService } from './type_materiels.service';
import { ClubsService } from '../clubs.service';
import { CreateTypeMaterielDto } from './dto/create-type_materiel.dto';
import { UpdateTypeMaterielDto } from './dto/update-type_materiel.dto';

@Controller('clubs/:club/type-materiels')
export class TypeMaterielsController {
  constructor(
    private readonly typeMaterielsService: TypeMaterielsService,
    private readonly clubsService: ClubsService,
  ) {}

  @Post()
  create(
    @Param('club') club: string,
    @Body() createTypeMaterielDto: CreateTypeMaterielDto,
  ) {
    const club_instance = this.clubsService.findBySlug(club);

    if (!club_instance) {
      throw new HttpException(
        {
          status: HttpStatus.NOT_FOUND,
          error: 'Club has not been found',
        },
        HttpStatus.NOT_FOUND,
      );
    }

    try {
      return this.typeMaterielsService.create(
        club_instance,
        createTypeMaterielDto,
      );
    } catch (err) {
      throw new HttpException(
        {
          status: HttpStatus.BAD_REQUEST,
          error: err.message,
        },
        HttpStatus.BAD_REQUEST,
      );
    }
  }

  @Get()
  findAll(@Param('club') club: string) {
    const club_instance = this.clubsService.findBySlug(club);

    if (!club_instance) {
      throw new HttpException(
        {
          status: HttpStatus.NOT_FOUND,
          error: 'Club has not been found',
        },
        HttpStatus.NOT_FOUND,
      );
    }

    return this.typeMaterielsService.findAll(club_instance);
  }

  @Get(':slug')
  findOne(@Param('club') club: string, @Param('slug') slug: string) {
    const club_instance = this.clubsService.findBySlug(club);

    if (!club_instance) {
      throw new HttpException(
        {
          status: HttpStatus.NOT_FOUND,
          error: 'Club has not been found',
        },
        HttpStatus.NOT_FOUND,
      );
    }

    const type_materiel = this.typeMaterielsService.findBySlug(
      club_instance,
      slug,
    );

    if (!type_materiel) {
      throw new HttpException(
        {
          status: HttpStatus.NOT_FOUND,
          error: 'Type material has not been found',
        },
        HttpStatus.NOT_FOUND,
      );
    }

    return type_materiel;
  }

  @Patch(':slug')
  update(
    @Param('club') club: string,
    @Param('slug') slug: string,
    @Body() updateTypeMaterielDto: UpdateTypeMaterielDto,
  ) {
    const club_instance = this.clubsService.findBySlug(club);

    if (!club_instance) {
      throw new HttpException(
        {
          status: HttpStatus.NOT_FOUND,
          error: 'Club has not been found',
        },
        HttpStatus.NOT_FOUND,
      );
    }

    try {
      return this.typeMaterielsService.update(
        slug,
        club_instance,
        updateTypeMaterielDto,
      );
    } catch (err) {
      throw new HttpException(
        {
          status: HttpStatus.BAD_REQUEST,
          error: err.message,
        },
        HttpStatus.BAD_REQUEST,
      );
    }
  }

  @Delete(':slug')
  remove(@Param('club') club: string, @Param('slug') slug: string) {
    const club_instance = this.clubsService.findBySlug(club);

    if (!club_instance) {
      throw new HttpException(
        {
          status: HttpStatus.NOT_FOUND,
          error: 'Club has not been found',
        },
        HttpStatus.NOT_FOUND,
      );
    }

    try {
      return this.typeMaterielsService.remove(club_instance, slug);
    } catch (err) {
      throw new HttpException(
        {
          status: HttpStatus.BAD_REQUEST,
          error: err.message,
        },
        HttpStatus.BAD_REQUEST,
      );
    }
  }
}

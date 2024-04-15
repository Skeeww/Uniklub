import {
  Body,
  Controller,
  DefaultValuePipe,
  Delete,
  Get,
  HttpException,
  HttpStatus,
  Param,
  ParseIntPipe,
  Patch,
  Post,
  Query,
} from '@nestjs/common';
import { CreateClubDTO } from './dto/create-club.dto';
import { UpdateClubDTO } from './dto/update-club.dto';
import { ClubsService } from './clubs.service';
import { Club } from './club.interface';
import slugify from 'slugify';

const slugifyOptions = {
  lower: true,
};

@Controller('clubs')
export class ClubsController {
  constructor(private clubsService: ClubsService) {}

  @Get()
  findAll(
    @Query(
      'limit',
      new ParseIntPipe({ optional: true }),
      new DefaultValuePipe(10),
    )
    limit: number,
  ) {
    return {
      limit: limit,
      clubs: this.clubsService.findAll().filter((_, idx) => idx < limit),
    };
  }

  @Get(':slug')
  findOne(@Param('slug') slug: string): Club {
    const club = this.clubsService.findBySlug(slug);

    if (!club) {
      throw new HttpException(
        {
          status: HttpStatus.NOT_FOUND,
          error: `club ${slug} has not been found`,
        },
        HttpStatus.NOT_FOUND,
      );
    }

    return club;
  }

  @Post()
  create(@Body() createClubDTO: CreateClubDTO): Club {
    const newClub: Club = {
      nom: createClubDTO.nom,
      slug: slugify(createClubDTO.nom, slugifyOptions),
    };

    try {
      this.clubsService.create(newClub);
    } catch (err) {
      throw new HttpException(
        {
          status: HttpStatus.BAD_REQUEST,
          error: err.message,
        },
        HttpStatus.BAD_REQUEST,
      );
    }

    return newClub;
  }

  @Patch(':slug')
  update(
    @Param('slug') slug: string,
    @Body() updateClubDTO: UpdateClubDTO,
  ): Club {
    const club = this.clubsService.findBySlug(slug);

    if (!club) {
      throw new HttpException(
        {
          status: HttpStatus.NOT_FOUND,
          error: `club ${slug} has not been found`,
        },
        HttpStatus.NOT_FOUND,
      );
    }

    const another_club = this.clubsService.findByName(updateClubDTO.nom);

    if (another_club) {
      throw new HttpException(
        {
          status: HttpStatus.BAD_REQUEST,
          error: `club ${another_club.nom} has already the same name`,
        },
        HttpStatus.BAD_REQUEST,
      );
    }

    club.nom = updateClubDTO.nom;
    return club;
  }

  @Delete(':slug')
  delete(@Param('slug') slug: string) {
    try {
      this.clubsService.delete(slug);
    } catch (err) {
      throw new HttpException(
        {
          status: HttpStatus.BAD_REQUEST,
          error: err.message,
        },
        HttpStatus.BAD_REQUEST,
      );
    }

    return HttpStatus.OK;
  }
}

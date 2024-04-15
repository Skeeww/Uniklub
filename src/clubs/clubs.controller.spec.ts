import { Test, TestingModule } from '@nestjs/testing';
import { randomUUID } from 'node:crypto';
import { ClubsController } from './clubs.controller';
import { ClubsService } from './clubs.service';

describe('ClubController', () => {
  let controller: ClubsController;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      controllers: [ClubsController],
      providers: [ClubsService],
    }).compile();

    controller = module.get<ClubsController>(ClubsController);
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
  });

  it('should return an empty clubs array', () => {
    const clubs = controller.findAll(10);
    expect(clubs.clubs.length).toStrictEqual(0);
  });

  it('should create a new club', () => {
    const club_name = randomUUID();

    const club = controller.create({
      nom: club_name,
    });

    expect(club).toHaveProperty('nom');
    expect(club).toHaveProperty('slug');
  });

  it('should not create a new club', () => {
    const club_name = randomUUID();
    const club_name_2 = club_name;

    const club_1 = controller.create({
      nom: club_name,
    });
    const club_2 = controller.create({
      nom: club_name_2,
    });

    expect(club_1).toHaveProperty('nom');
    expect(club_1).toHaveProperty('slug');

    expect(club_2).toHaveProperty('error');
    expect(club_2).not.toHaveProperty('nom');
    expect(club_2).not.toHaveProperty('slug');
  });
});

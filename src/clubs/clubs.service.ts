import { Injectable } from '@nestjs/common';
import { Club } from './club.interface';

@Injectable()
export class ClubsService {
  private readonly clubs: Club[] = [];

  create(club: Club) {
    if (this.findByName(club.nom) || this.findBySlug(club.slug)) {
      throw new Error('this club already exists');
    }
    this.clubs.push(club);
  }

  findAll(): Club[] {
    return this.clubs;
  }

  /**
   * Find a club by its slug name
   * @param slug Sluggified name of the club
   * @returns A club instance or undefined if not found
   */
  findBySlug(slug: string): Club | undefined {
    const club = this.clubs.find((c) => c.slug === slug);
    return club;
  }

  findByName(name: string): Club | undefined {
    const club = this.clubs.find((c) => c.nom === name);
    return club;
  }

  delete(slug: string) {
    const club = this.findBySlug(slug);

    if (!club) {
      throw new Error('this club does not exist');
    }

    this.clubs.splice(this.clubs.indexOf(club), 1);
  }
}

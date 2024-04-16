import { Injectable } from '@nestjs/common';
import { CreateTypeMaterielDto } from './dto/create-type_materiel.dto';
import { UpdateTypeMaterielDto } from './dto/update-type_materiel.dto';
import { TypeMateriel } from './type_materiel.interface';
import { Club } from '../club.interface';
import slugify from 'slugify';

const slugifyOptions = {
  lower: true,
};

@Injectable()
export class TypeMaterielsService {
  private readonly type_materiels: TypeMateriel[] = [];

  create(
    club: Club,
    createTypeMaterielDto: CreateTypeMaterielDto,
  ): TypeMateriel {
    const type_materiel: TypeMateriel = {
      ...createTypeMaterielDto,
      slug: slugify(createTypeMaterielDto.nom, slugifyOptions),
      club: club,
    };

    if (
      this.findByName(club, type_materiel.nom) ||
      this.findBySlug(club, type_materiel.slug)
    ) {
      throw new Error('Type material already exists');
    }

    this.type_materiels.push(type_materiel);
    return type_materiel;
  }

  findByName(club: Club, nom: string): TypeMateriel | undefined {
    const type_materiel = this.type_materiels.find(
      (t) => t.club.nom === club.nom && t.nom === nom,
    );
    return type_materiel;
  }

  findBySlug(club: Club, slug: string): TypeMateriel | undefined {
    const type_materiel = this.type_materiels.find(
      (t) => t.club.nom === club.nom && t.slug === slug,
    );
    return type_materiel;
  }

  findAll(club: Club) {
    return this.type_materiels.filter((t) => t.club.nom === club.nom);
  }

  update(
    slug: string,
    club: Club,
    updateTypeMaterielDto: UpdateTypeMaterielDto,
  ) {
    const type_materiel = this.findBySlug(club, slug);
    if (!type_materiel) {
      throw new Error('The material type has not been found');
    }

    const new_type_material: TypeMateriel = {
      ...updateTypeMaterielDto,
      slug: slugify(updateTypeMaterielDto.nom, slugifyOptions),
      club: club,
    };
    if (
      this.findByName(club, new_type_material.nom) ||
      this.findBySlug(club, new_type_material.slug)
    ) {
      throw new Error('The new material type already exists');
    }

    type_materiel.nom = new_type_material.nom;
    type_materiel.slug = new_type_material.slug;
    return type_materiel;
  }

  remove(club: Club, slug: string) {
    const type_materiel = this.findBySlug(club, slug);

    if (!type_materiel) {
      throw new Error('The material type has not been found');
    }

    this.type_materiels.splice(this.type_materiels.indexOf(type_materiel), 1);
  }
}

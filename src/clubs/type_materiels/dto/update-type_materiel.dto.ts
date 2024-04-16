import { PartialType } from '@nestjs/mapped-types';
import { CreateTypeMaterielDto } from './create-type_materiel.dto';

export class UpdateTypeMaterielDto extends PartialType(CreateTypeMaterielDto) {
  nom: string;
}

import { Entity, Column, PrimaryGeneratedColumn, Generated } from 'typeorm';

@Entity
export class user_info{

    @PrimaryGeneratedColumn('uuid')
    UUID: string

}
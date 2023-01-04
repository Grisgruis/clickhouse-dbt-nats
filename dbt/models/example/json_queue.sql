{{
    config(
        materialized='incremental'
    )
}}

/* Only publish data which does not exist yet */
select json from data where json not in (select raw_json from json_data)
# Changelog

All notable changes to this project will be documented in this file. See [standard-version](https://github.com/conventional-changelog/standard-version) for commit guidelines.

## 0.1.0 (2025-06-04)


### Build System

* **docker:** Add mount volume for database ([342a6b4](https://github.com/mushoffa/payd/commit/342a6b49d7c5cf68d65cf4f14963110dc04d833f))
* **docker:** Initialize docker-compose file on database container ([4f0d14f](https://github.com/mushoffa/payd/commit/4f0d14f5044ebd5be96c4dd608dff12fb261c018))


### Database

* **migration:** Create shifts and shift_requests table ([0a88dcb](https://github.com/mushoffa/payd/commit/0a88dcbb145a77bb4bc2a14e8689ddd9e196e314))
* **migration:** Update shift_requests table column and index ([2288975](https://github.com/mushoffa/payd/commit/228897552121ef60953d23ccd71ab9a49a19b49f))


### Documentation

* **README:** Add additional assumption(s) of business requirements ([3bacad6](https://github.com/mushoffa/payd/commit/3bacad66fb860b381238d5a38625e51bf6ff44db))


### Backend

* Bump dependency library ([154fc57](https://github.com/mushoffa/payd/commit/154fc576ef843310065ddedd77f61db24204ca1b))
* **domain:** Add date format regex validation ([54f7cd3](https://github.com/mushoffa/payd/commit/54f7cd38137cdd8e083963a9ac1c3f932aa6f34b))
* **domain:** Add date regex unit test ([ec0e470](https://github.com/mushoffa/payd/commit/ec0e470b40444460834fb357c6c081e9cdb146c2))
* **domain:** Add domain error data structure as shared kernel ([e5d59a5](https://github.com/mushoffa/payd/commit/e5d59a5ebb0ef4d3cd10960506fe0a5448730166))
* **domain:** Add shift repository interface ([c99306c](https://github.com/mushoffa/payd/commit/c99306ccd33096658882c33b4fe1946224aa3146))
* **domain:** Add time format regex validation ([1fa45d1](https://github.com/mushoffa/payd/commit/1fa45d1158564ff21cc46b5f47a3af0a508f9572))
* **domain:** Add time regex unit test ([63fc239](https://github.com/mushoffa/payd/commit/63fc239c4f33bd1576d43d4d7e4d8a09ee16a092))
* **entity:** Add new Shift data structure ([e6f59d1](https://github.com/mushoffa/payd/commit/e6f59d1ed974f1eab247304be07393fbf66fe101))
* **entity:** Add shift data structure ([cee8a72](https://github.com/mushoffa/payd/commit/cee8a725dc5e1b9fc54cb252f0367b3b83e37823))
* **infrastructure:** Add AddRoute function on server struct ([634d372](https://github.com/mushoffa/payd/commit/634d372ffa37c98502d45c1f14d952a75c63733e))
* **infrastructure:** Add custom http request validator data structure ([73d9b9b](https://github.com/mushoffa/payd/commit/73d9b9be26064318151e2594067b3c4182500641))
* **infrastructure:** Add message field error response ([784259f](https://github.com/mushoffa/payd/commit/784259fb7503f3a923dc4faed27116ee24d7dcad))
* **infrastructure:** Add postgres client ada database service generic interface ([b37ada4](https://github.com/mushoffa/payd/commit/b37ada473e32787431d34f9710a3b7fcad40fe1a))
* **infrastructure:** Encapsulate infrastructure instances into single struct ([d7d9d10](https://github.com/mushoffa/payd/commit/d7d9d105eb215dac41c46b59422ea4b10fba6c3b))
* **infrastructure:** Integrate fiber handler with custom validator ([ec50f53](https://github.com/mushoffa/payd/commit/ec50f53bb420226523eb5b1a34c5b7b67ca8b3ae))
* **infrastructure:** Refactor http server package ([34a4562](https://github.com/mushoffa/payd/commit/34a4562b3ba1c372feb19dd5840dc038a327b46a))
* **Shift:** Add basic CRUD function interface ([e0c0d49](https://github.com/mushoffa/payd/commit/e0c0d49f4e2d11c3abe917aff267d38b62aa61f5))
* **shift:** Add database conn interface ([482b603](https://github.com/mushoffa/payd/commit/482b603e19dfe7cd29f33280c5530d6676c1daf6))
* **shift:** Add ShiftDate and ShiftTime error types ([6c9d99b](https://github.com/mushoffa/payd/commit/6c9d99b3eb1ea0643b441b7186659d8f9c48c3c6))
* **shift:** Refactor error return type ([357bfbe](https://github.com/mushoffa/payd/commit/357bfbec6314efd43ed09d845ed3b2dd9ca64984))
* **Shift:** Refactor to vertical slice architecture ([b5983ca](https://github.com/mushoffa/payd/commit/b5983ca36d22762c603db0a40750d4402e3de9fa))
* **test:** Add shift_date, shift_time unit test ([ccd57d5](https://github.com/mushoffa/payd/commit/ccd57d5aebaabf232bc3b9d745c7f9b14cf13016))
* **test:** Add ShiftDate backdate test case ([e21c46f](https://github.com/mushoffa/payd/commit/e21c46fee31c7a6a1d635d5cee53736027d4d883))
* **test:** Add ShiftTime past hour & minute test case ([1159953](https://github.com/mushoffa/payd/commit/11599539dbd669e0e19f474f5347f073e30513e8))
* Tie the string up ([e3dea66](https://github.com/mushoffa/payd/commit/e3dea6680758cb91073abdf3ba2477c5a541e8c3))
* **valueobject:** Add predefined shift roles ([4e0b4f4](https://github.com/mushoffa/payd/commit/4e0b4f46ba0aa767f01465ddf680f7595b9a5063))
* **valueobject:** Add shift_date and shift_time ([bfc902c](https://github.com/mushoffa/payd/commit/bfc902cd0baca3b4f52a106d64e2f46cb30da50d))
* **valueobject:** Add shift_date format validation ([583e19b](https://github.com/mushoffa/payd/commit/583e19bf8c535a68533502863b5552b7b10b8bd3))
* **valueobject:** Add shift_time format validation ([933f253](https://github.com/mushoffa/payd/commit/933f253a5a0a5532446206fb4d156a89b5409cf4))
* **valueobject:** Add ShiftDate factory method and backdate validation ([6ee2b67](https://github.com/mushoffa/payd/commit/6ee2b674746892184e68240203317f76407f26ef))
* **valueobject:** Add ShiftTime factory method and formatted string ([a375ef8](https://github.com/mushoffa/payd/commit/a375ef883a938586c28a4002bf36e485f3b1b13f))

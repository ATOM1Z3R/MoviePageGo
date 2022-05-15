import { CreateCommentModel, ReadCommentModel } from "../model/comment";

export class CommentApi{
    apiAddress = 'http://localhost:8080/comment/';

    async createCommentApiCallAsync(comment: CreateCommentModel, token: string): Promise<ReadCommentModel[]>{
        const request = {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'token': token,
            },
            body: JSON.stringify(comment),
        };
        
        const jsonString = await fetch(`${this.apiAddress}create`, request).then(async (response) => {
            const body = (await response.body?.getReader().read())?.value
            return new TextDecoder().decode(body);
        });
        let jsonObjList = JSON.parse(jsonString);
        return jsonObjList.map((x: any)=><ReadCommentModel>x);
    }
}

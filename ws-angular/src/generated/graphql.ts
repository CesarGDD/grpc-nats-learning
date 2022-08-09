import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
};

export type Blog = {
  __typename?: 'Blog';
  content: Scalars['String'];
  id: Scalars['Int'];
  username: Scalars['String'];
};

export type Mutation = {
  __typename?: 'Mutation';
  createBlog: Blog;
  deleteBlog: Blog;
};


export type MutationCreateBlogArgs = {
  input: NewBlog;
};


export type MutationDeleteBlogArgs = {
  id: Scalars['Int'];
};

export type NewBlog = {
  content: Scalars['String'];
  username: Scalars['String'];
};

export type Query = {
  __typename?: 'Query';
  blog: Blog;
  blogs: Array<Blog>;
};


export type QueryBlogArgs = {
  id: Scalars['Int'];
};

export type Subscription = {
  __typename?: 'Subscription';
  blog: Blog;
};

export type BlogsQueryVariables = Exact<{ [key: string]: never; }>;


export type BlogsQuery = { __typename?: 'Query', blogs: Array<{ __typename?: 'Blog', id: number, username: string, content: string }> };

export type BlogQueryVariables = Exact<{
  blogId: Scalars['Int'];
}>;


export type BlogQuery = { __typename?: 'Query', blog: { __typename?: 'Blog', id: number, username: string, content: string } };

export type CreateBlogMutationVariables = Exact<{
  input: NewBlog;
}>;


export type CreateBlogMutation = { __typename?: 'Mutation', createBlog: { __typename?: 'Blog', id: number, username: string, content: string } };

export type DeleteBlogMutationVariables = Exact<{
  deleteBlogId: Scalars['Int'];
}>;


export type DeleteBlogMutation = { __typename?: 'Mutation', deleteBlog: { __typename?: 'Blog', id: number, username: string, content: string } };

export type BlogSubSubscriptionVariables = Exact<{ [key: string]: never; }>;


export type BlogSubSubscription = { __typename?: 'Subscription', blog: { __typename?: 'Blog', id: number, username: string, content: string } };

export const BlogsDocument = gql`
    query Blogs {
  blogs {
    id
    username
    content
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class BlogsGQL extends Apollo.Query<BlogsQuery, BlogsQueryVariables> {
    // @ts-ignore
    document = BlogsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const BlogDocument = gql`
    query Blog($blogId: Int!) {
  blog(id: $blogId) {
    id
    username
    content
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class BlogGQL extends Apollo.Query<BlogQuery, BlogQueryVariables> {
        // @ts-ignore
    document = BlogDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const CreateBlogDocument = gql`
    mutation CreateBlog($input: NewBlog!) {
  createBlog(input: $input) {
    id
    username
    content
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class CreateBlogGQL extends Apollo.Mutation<CreateBlogMutation, CreateBlogMutationVariables> {
        // @ts-ignore
    document = CreateBlogDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const DeleteBlogDocument = gql`
    mutation DeleteBlog($deleteBlogId: Int!) {
  deleteBlog(id: $deleteBlogId) {
    id
    username
    content
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class DeleteBlogGQL extends Apollo.Mutation<DeleteBlogMutation, DeleteBlogMutationVariables> {
        // @ts-ignore
    document = DeleteBlogDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const BlogSubDocument = gql`
    subscription BlogSub {
  blog {
    id
    username
    content
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class BlogSubGQL extends Apollo.Subscription<BlogSubSubscription, BlogSubSubscriptionVariables> {
        // @ts-ignore
    document = BlogSubDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }